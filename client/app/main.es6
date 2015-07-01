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
//  Set the layout and Router
//
app.config(($stateProvider, $locationProvider, $urlRouteProvider) => {
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

    //  Unmatched URL state
    $urlRouterProvider.otherwise("/")
})

//
//  Restanuglar setup
//
app.config((RestangularProvider) => {
    RestangularProvider.setBaseUrl('/api/v1')

    RestangularProvider.setResponseExtractor((response, operation, what, url) => {
            //  This is a get for a list
            if (operation === "getList") {
                //  Here we're returning an Array which has one special property metadata with our extra information
                return response.data
            } else  {
                //  This is an element
                return response.data
            }
        })
})

app.run(($rootScope, $state, $location, AuthService) => {
    $rootScope.$on("$stateChangeStart", (event, toState, toParams, fromState, fromParams) => {
        if (toState.authenticate && !AuthService.isAuthenticated()) {
            //  If the user isn't authenticated, redirect to the login page with a good "next" URL
            let href = $state.href(toState, toParams)
            $state.transitionTo("app.auth.login", { next: $location.path() })
            event.preventDefault()
        }
    })
})
