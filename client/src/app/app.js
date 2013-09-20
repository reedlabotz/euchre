angular.module('app', [
    'game',
    'templates.app']);

angular.module('app').config(['$routeProvider', function ($routeProvider) {
    $routeProvider.otherwise({redirectTo:'/welcome'});
}]);


angular.module('app').controller('AppCtrl', ['$scope', function($scope) {
}]);