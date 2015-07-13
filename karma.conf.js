// Karma configuration
// Generated on Sat Jul 11 2015 20:26:35 GMT-0700 (PDT)

module.exports = function(config) {
  config.set({

    // base path that will be used to resolve all patterns (eg. files, exclude)
    basePath: '',

    // frameworks to use
    // available frameworks: https://npmjs.org/browse/keyword/karma-adapter
    // frameworks: ['browserify', 'mocha', 'chai'],
    frameworks: ['browserify', 'jasmine'],


    // list of files / patterns to load in the browser
    files: [
        'client/components/angular/angular.js',
        'client/components/underscore/underscore.js',
        'client/components/angular-ui-router/release/angular-ui-router.js',
        'client/components/angular-cookies/angular-cookies.js',
        'client/components/restangular/dist/restangular.js',
        'client/components/angular-bootstrap/ui-bootstrap.js',
        'client/components/angular-bootstrap/ui-bootstrap-tpls.js',
        'client/components/angular-mocks/angular-mocks.js',

        'client/app/**/*.es6',
        'client/test/**/*.spec.es6',
        'client/test/**/*.spec.js'
    ],


    // list of files to exclude
    exclude: [
      '**/*.swp'
    ],


    // preprocess matching files before serving them to the browser
    // available preprocessors: https://npmjs.org/browse/keyword/karma-preprocessor
    preprocessors: {
        'client/**/*.es6': ['browserify']
    },


    // test results reporter to use
    // possible values: 'dots', 'progress'
    // available reporters: https://npmjs.org/browse/keyword/karma-reporter
    reporters: ['progress'],


    // web server port
    port: 9876,


    // enable / disable colors in the output (reporters and logs)
    colors: true,


    // level of logging
    // possible values: config.LOG_DISABLE || config.LOG_ERROR || config.LOG_WARN || config.LOG_INFO || config.LOG_DEBUG
    // logLevel: config.LOG_INFO,
    logLevel: config.LOG_DEBUG,


    // enable / disable watching file and executing tests whenever any file changes
    autoWatch: true,


    // start these browsers
    // available browser launchers: https://npmjs.org/browse/keyword/karma-launcher
    browsers: ['Chrome'],

    //
    browserify: {
        transform: [
            ["babelify", {
                loose: "all",
                sourceMap: true,
            }], "partialify",
        ],
        extensions: ['.es6'],
        debug: true,
        comments: true
    },


    // Continuous Integration mode
    // if true, Karma captures browsers, runs the tests and exits
    singleRun: false
  })
}
