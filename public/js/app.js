
tetrisApp = angular.module("tetrisApp" , [])
  // .config(function($interpolateProvider) {
  // 	  $interpolateProvider.startSymbol('[[');
  // 	  $interpolateProvider.endSymbol(']]');
  // })
tetrisApp.controller("HomepageCtrl",["$scope", function($scope) {
    $scope.game_over = true;
    $scope.loading = true;
    $scope.teams = 1;


    $scope.pusher = new Pusher('80f71c71ecfd0ce866eb');
    $scope.channel = $scope.pusher.subscribe('my_channel');


    $scope.channel.bind('my_event', function(data) {
    	$scope.gameData = data.cells;
      $scope.game_over = data.game_over;
      // $scope.gameNextPiece = data.next_piece;

      if(data.game_over) {
        alert('GAME OVER');
        return
      }


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
      //   piece_size = ($scope.gameNextPiece.Shape[0]).length;

      //   var parts = [];

      //   for (var i = 0; i < piece_size; i=i+2)
      //   {

      //       parts.push(parseInt($scope.gameNextPiece.Shape[0][i] + $scope.gameNextPiece.Shape[0][i]));
      //   }


      // for (var i = 0; i < piece_size/2; i++)
      //   {

      //   console.log("first byte : " + bytes[0]);
      //   }
      //   // console.log("first byte : " + bytes[0]);


        // var canvas2 = document.getElementById('myCanvas2');
        // var context2 = canvas2.getContext('2d');

        // canvas.width = canvas.width;

        // piece_size = ($scope.gameNextPiece.Shape[0]).length/2;
        // console.log($scope.gameNextPiece)
        // console.log(piece_size);

        // for(i=0; i<piece_size; i++){
        //   for(j=0; i<piece_size; j++ ) {

        //     // if ($scope.gameData[i][j] != 0){
        //       context.fillStyle=colors[$scope.gameData[i][j]];
        //       context.strokeStyle="white";
        //       context.strokeRect( ((canvas2.width/piece_size)*j) ,((canvas2.height/piece_size)*i),canvas2.width/piece_size, canvas2.height/piece_size);

        //       context.fillRect( ((canvas2.width/piece_size)*j) ,((canvas2.height/piece_size)*i),canvas2.width/piece_size, canvas2.height/piece_size);
 
            // }





        //   }

        // }

    



		


	  }

    	$scope.drawPieces();



    });



}]);



  // $scope.drawPieces = function () {
  //       var canvas = document.getElementById('myCanvas');
  //       var context = canvas.getContext('2d');

 

  // }

