
tetrisApp = angular.module("tetrisApp" , [])
  // .config(function($interpolateProvider) {
  // 	  $interpolateProvider.startSymbol('[[');
  // 	  $interpolateProvider.endSymbol(']]');
  // })
tetrisApp.controller("HomepageCtrl",["$scope", function($scope) {

    $scope.loading = true;


    $scope.pusher = new Pusher('80f71c71ecfd0ce866eb');
    $scope.channel = $scope.pusher.subscribe('my_channel');

    // console.log("test test ");

    $scope.channel.bind('my_event', function(data) {
      console.log(data.message);
    });



}]);



  // $scope.drawPieces = function () {
  //       var canvas = document.getElementById('myCanvas');
  //       var context = canvas.getContext('2d');

 

  // }

