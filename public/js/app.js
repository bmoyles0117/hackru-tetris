
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
    	$scope.gameData = data;
      	console.log("hit");


      $scope.drawPieces = function () {
        var canvas = document.getElementById('myCanvas');
        var context = canvas.getContext('2d');

        canvas.width = canvas.width;

        for(i=0; i< 22; i++){
        	for(j=0; j<12; j++ ) {
        		if ($scope.gameData[j][i]){
	        		context.rect( (j+1)*20,(i+1)*20,20,20);
					context.fillStyle="blue";
					context.fill();
				}
        	}
        }

		


	  }

    	$scope.drawPieces();



    });



}]);



  // $scope.drawPieces = function () {
  //       var canvas = document.getElementById('myCanvas');
  //       var context = canvas.getContext('2d');

 

  // }

