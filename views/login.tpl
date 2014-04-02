<!DOCTYPE html>

<html>
  	<head>
    	<title>Beego</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	
		<style type="text/css">
			 .container{
			 	background-color: #cccccc;
			 }
       .error{
        color: red;
       }
		</style>
	</head>
  	<body>
  		 <div class="container">
  		 	<form method="POST">
          <div class="error"> {{.Error}} </div>
  		 		<div>
  		 			<lable>username:</lable><input type="text" name="username" size="20"/>
  		 		</div>
  		 		<div>
  		 			<lable>password:</lable><input type="password" name="password" size="20"/>
  		 		</div>
  		 		<div>
  		 			<input type="submit" size="20" value="Login"></input>
  		 		</div>
  		 	</form>
  		 </div>
	</body>
</html>
