var app = angular.module('iqvine', ['restangular', 'ui.router', 'ngCookies', 'ui.bootstrap'])

angular.element(document).ready(() => angular.bootstrap(document, ['iqvine']))

import AuthService from './services/auth'
import { AuthRegister } from './controllers/auth'
import { TodoRegister } from './controllers/todo'
import { MainRegister } from './controllers/main'

// require('./services/auth')
// require('./controllers/main')
// require('./controllers/todo')
// require('./controllers/auth')

app.service('AuthService', AuthService)

AuthRegister(app)
TodoRegister(app)
MainRegister(app)

//
//  Set the layout
//
app.config(($stateProvider, $locationProvider) => {
    $stateProvider
        .state('app', {
            url: '',
            abstract: true,
            views: {
                container: {
                    template: require('../partials/layout.html')
                }, 
                footer: {
                    template: require('../partials/footer.html')
                },
                header: {
                    controller: 'HeaderController',
                    template: require('../partials/header.html')
                }
            }
        })
    $locationProvider.html5Mode({
            enabled: true,
            requireBase: false
        })
})

//
//  Restanuglar setup
//
app.config((RestangularProvider, $stateProvider, $urlRouterProvider) => {
    RestangularProvider.setBaseUrl('/api/v1')

    RestangularProvider.setResponseExtractor((response, operation, what, url) => {
            var newResponse

            //  This is a get for a list
            if (operation === "getList") {
                //  Here we're returning an Array which has one special property metadata with our extra information
                newResponse = response.data
            } else  {
                //  This is an element
                newResponse = response.data
            }
            return newResponse
        })

    //  Unmatched URL state
    $urlRouterProvider.otherwise("/")
})

app.run(($rootScope, $state, $location, AuthService) => {
    $rootScope.$on("$stateChangeStart", (event, toState, toParams, fromState, fromParams) => {
        if (toState.authenticate && !AuthService.isAuthenticated()) {
            //  User isn’t authenticated
            let href = $state.href(toState, toParams)
            $state.transitionTo("app.auth.login", { next: $location.path() })
            event.preventDefault()
        }
    })
})
