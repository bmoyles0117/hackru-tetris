
tetrisApp = angular.module("tetrisApp" , [])
  // .config(function($interpolateProvider) {
  // 	  $interpolateProvider.startSymbol('[[');
  // 	  $interpolateProvider.endSymbol(']]');
  // })
tetrisApp.controller("HomepageCtrl",["$scope", function($scope) {

    $scope.loading = true;


    $scope.pusher = new Pusher('80f71c71ecfd0ce866eb');
    $scope.channel = $scope.pusher.subscribe('my_channel');


    $scope.channel.bind('my_event', function(data) {
    	$scope.gameData = data.board;



      $scope.drawPieces = function () {
        var canvas = document.getElementById('myCanvas');
        var context = canvas.getContext('2d');

        canvas.width = canvas.width;


       colors = {
          73 : "blue",
          74 : "green",
          76 : "yellow",
          79 : "red",
          83 : "cyan",
          84 : "black",
          90 : "gray"
       };

        for(i=0; i<22; i++){
        	for(j=0; j<12; j++ ) {

            if ($scope.gameData[i][j] != 0){
    					context.fillStyle=colors[$scope.gameData[i][j]];
              context.strokeStyle="white";
              context.strokeRect( ((canvas.width/12)*j) ,((canvas.height/22)*i),canvas.width/12, canvas.height/22);

              context.fillRect( ((canvas.width/12)*j) ,((canvas.height/22)*i),canvas.width/12, canvas.height/22);
 
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

