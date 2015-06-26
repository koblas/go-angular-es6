import app from '../app'

class AuthService {
    /*@ngInject*/
    constructor(Restangular, $q, $cookies, $state) {
        this.$cookies = $cookies
        this.$q = $q

        this.Auth = Restangular.all('auth')
        this.auth_cookie = 'user_auth'
        this.authenticated = false
        this.name = null

        if ($cookies[this.auth_cookie])  {
            this.authenticated = true

            this.Auth.post({token:$cookies[this.auth_cookie]})
                .then((auth) => {
                    if (!auth && !auth.token)  {
                        //  If the token is "bad" e.g. you're no longer a valid user, cleanup
                        delete $cookies[this.auth_cookie]
                        this.authenticated = false
                        $state.transitionTo("index")
                    }
                })
        }
    }

    isAuthenticated() {
        return this.authenticated
    }

    //
    //  Login a user by email + password - return a promise
    // 
    //  TODO - better error messages on the result
    //
    login(email, password) {
        var deferred = this.$q.defer()

        this.Auth.post({email:email, password:password})
            .then((auth) => {
                if (auth.token) {
                    this.authenticated = true

                    this.$cookies[this.auth_cookie] = auth.token

                    deferred.resolve("ok")
                } else {
                    deferred.reject("unknown")
                }
            }).catch((err) => {
                deferred.reject(err.data.emsg)
            })

        return deferred.promise
    }

    logout() {
        this.authenticated = false
        delete this.$cookies[this.auth_cookie]
    }

    register(email, password, params) {
        var deferred = this.$q.defer()

        this.Auth.post({email:email, password:password, params:params}, {register:true})
            .then((auth) => {
                if (auth.token) {
                    this.authenticated = true

                    this.$cookies[this.auth_cookie] = auth.token
                    deferred.resolve("ok")
                } else {
                    deferred.reject("unknown")
                }
            }).catch((err) =>
                deferred.reject(err.data.emsg)
            )

        return deferred.promise
    }
}

app.service('AuthService', AuthService)
