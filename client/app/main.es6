/*
**  Good Reference
**       http://www.michaelbromley.co.uk/blog/350/exploring-es6-classes-in-angularjs-1-x
**
*/

import app from './app'

import './services/auth'
import './controllers/auth'
import './controllers/todo'
import './controllers/main'

//
//  Bootstrap in the UI
//
angular.element(document).ready(() => angular.bootstrap(document, [app.name]))

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
