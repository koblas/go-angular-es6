package app

import (
    "strings"
    "strconv"
    "errors"
    "fmt"
    "time"
    "encoding/base64"
    "crypto/hmac"
    "crypto/sha256"
)

func format_field(s string) string {
    return fmt.Sprintf("%d:%s", len(s), s)
}

func decodeField(s string) (string, string, error) {
    parts := strings.SplitN(s, ":", 2)

    if len(parts) != 2 {
        return "", "", errors.New("Missing lensep")
    }

    length, err := strconv.ParseInt(parts[0], 10, 64)

    if err != nil {
        return "", "", errors.New("Bad length")
    }
    if parts[1][length] != '|' {
        return "", "", errors.New("Not at separator")
    }

    return parts[1][0:length], parts[1][length+1:], nil
}

// The v2 format consists of a version number and a series of
// length-prefixed fields "%d:%s", the last of which is a
// signature, all separated by pipes.  All numbers are in
// decimal format with no leading zeros.  The signature is an
// HMAC-SHA256 of the whole string up to that point, including
// the final pipe.
//
// The fields are:
// - format version (i.e. 2; no length prefix)
// - key version (integer, default is 0)
// - timestamp (integer seconds since epoch)
// - name (not encoded; assumed to be ~alphanumeric)
// - value (base64-encoded)
// - signature (hex-encoded; no length prefix)

func CreateSignedValue(secret string, name string, value string, clock *time.Time) string {
    var timestamp string

    if clock == nil {
        timestamp = fmt.Sprintf("%d", time.Now().Unix())
    } else {
        timestamp = fmt.Sprintf("%d", clock.Unix())
    }

    valueStr := base64.StdEncoding.EncodeToString([]byte(value))

    to_sign := strings.Join([]string{"2",
                format_field("0"),
                format_field(timestamp),
                format_field(name),
                format_field(valueStr),
                ""}, "|")

    signature := createSignatureV2(secret, to_sign)
    return to_sign + signature
}

func DecodeSignedValue(secret string, name string, value string, clock *time.Time) (string, error) {
    /*
    var timestamp string

    if clock == nil {
        timestamp = fmt.Sprintf("%d", time.Now().Unix())
    } else {
        timestamp = fmt.Sprintf("%d", clock.Unix())
    }
    */

    var (
        err error
        // key_version string
        // timestamp_field string
        name_field string
        value_field string
        passed_sig string
    )

    rest := value[2:]
    _, rest, err = decodeField(rest)        // key_version
    if err != nil {
        return "", err
    }
    _, rest, err = decodeField(rest)        // timestamp_field
    if err != nil {
        return "", err
    }
    name_field, rest, err = decodeField(rest)
    if err != nil {
        return "", err
    }
    value_field, passed_sig, err = decodeField(rest)
    if err != nil {
        return "", err
    }

    if name_field != name {
        return "", errors.New("Bad field name")
    }

    // TODO: timestamp < clock() - max_age_days

    signed_string := value[0:len(value)-len(passed_sig)]
    expected_sig := createSignatureV2(secret, signed_string)

    if passed_sig != expected_sig {
        return "", errors.New("Signature mismatch")
    }

    data, err := base64.StdEncoding.DecodeString(value_field)
    if err != nil {
        return "", errors.New("Base64 decode failed")
    }

    return string(data), nil
}

func createSignatureV2(secret string, message string) string {
    key := []byte(secret)
    h := hmac.New(sha256.New, key)
    h.Write([]byte(message))
    return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
