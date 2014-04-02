<!DOCTYPE html>

<html>
  	<head>
    	<title>Register</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	
		<style type="text/css">
			 .container{
			 	background-color: #cccccc;
			 }
       .error{
        color: red;
       }
       .success{
        color:green;
       }
		</style>
	</head>
  	<body>
  		 <div class="container">
  		 	<form method="POST">
          <div class="error"> {{.Error}} </div>
          <div class="success"> {{.Success}} </div>
  		 		<div>
  		 			<lable>username:</lable><input type="text" name="username" size="20"/>
  		 		</div>
  		 		<div>
  		 			<lable>password:</lable><input type="password" name="password" size="20"/>
  		 		</div>
          <div>
            <lable>Email:</lable><input type="text" name="email" size="20"/>
          </div>
          <div>
            <lable>Sex:</lable><input type="checkbox" name="sex"/>
          </div>
  		 		<div>
  		 			<input type="submit" size="20" value="Register"></input>
            <a href="/login">toLogin</a>
  		 		</div>
  		 	</form>
  		 </div>
	</body>
</html>
