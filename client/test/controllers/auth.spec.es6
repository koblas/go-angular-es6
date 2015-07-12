describe("Register Controller:", function() {
    beforeEach(angular.mock.module('app'));

    let registerController, scope;

    beforeEach(angular.mock.inject(($rootScope, $controller, _$location_) => {
        scope = $rootScope.$new();

        registerController = $controller("RegisterController", {    
                $location: _$location_,
                AuthService: null,
                $stateParams: { next: "/" },
                $scope: scope, 
            });
    }));

    //
    //  Start Tests
    //

    it('initial state', () => {
        expect(registerController.username).toEqual('')
    })
});
