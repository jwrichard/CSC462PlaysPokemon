<?php
	if(empty($_POST['move'])){
		echo "400";
	} else {
		// Generate a random color for the clients messages
		$clientNum = ip2long($_SERVER['REMOTE_ADDR']);
		$clientNum2 = (int)strrev($clientNum.'');
		$magicNumber = $clientNum*$clientNum2;
		$rgbstring = 'rgb('. ($clientNum % 256) .','. ($clientNum2 % 256) .','. ($magicNumber % 256) .')';
		$str = "<span style='color:$rgbstring;'>[".date("g:i:s")."]</span> ".htmlentities($_POST['move']).'<br>';
		
		
		$file = fopen("chat.txt", "a") or die("Unable to open file!");
		if (flock($file, LOCK_EX)) { // do an exclusive lock
			fwrite($file, $str);
			flock($file, LOCK_UN); // release the lock
		} else {
			echo "Couldn't lock the file !";
		}
		fclose($file);
		
		$file = fopen("/home/ubuntu/chat/chat.txt", "a") or die("Unable to open file!");
		$str = htmlentities($_POST['move']).PHP_EOL;
		if (flock($file, LOCK_EX)) { // do an exclusive lock
			fwrite($file, $str);
			flock($file, LOCK_UN); // release the lock
		} else {
			echo "Couldn't lock the file !";
		}
		fclose($file);
		
		echo "200";
	}
?>
