<?php
	if(empty($_POST['move'])){
		echo "400";
	} else {
		// Generate a random color for the clients messages
		$clientNum = ip2long($_SERVER['REMOTE_ADDR']);
		$rgbstring = 'rgb('. $clientNum%256 .','. (255-($clientNum%256)) .','. $clientNum%256 .')';
		$str = "<span style='color:$rgbstring;'>[".date("g:i:s")."]</span> ".htmlentities($_POST['move']).'<br>';
		$file = fopen("chat.txt", "a") or die("Unable to open file!");
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
