import app from '../app'

app.config(($stateProvider) => {
    $stateProvider
        .state('app.index_', {
            url: '',
            views: {
                'content@app': {
                    template: require('../../partials/index.html'),
                    controller: "IndexController",
                    authenticate: false
                }
            }
        });
    $stateProvider
        .state('app.index', {
            url: '/',
            views: {
                'content@app': {
                    template: require('../../partials/index.html'),
                    controller: "IndexController",
                    authenticate: false
                }
            }
        });
})

app.controller('IndexController', class IndexController {
    /*@ngInject*/
    constructor(AuthService, $state) {
        if (AuthService.isAuthenticated())
            $state.go('app.dashboard')
    }
})

app.controller('HeaderController', class HeaderController {
    /*@ngInject*/
    constructor(AuthService, $location) {
        this.auth = AuthService;
        this.$location = $location
    }

    logout() {
        this.auth.logout();
        this.$location.path('/');
    }
})
