<?php
    function flatten($array) { 
        if (!is_array($array)) { 
        return false; 
        } 
        $result = array(); 
        foreach ($array as $key => $value) { 
        if (is_array($value)) { 
            $result = array_merge($result, flatten($value)); 
        } else { 
            $result = array_merge($result, array($key => $value));
        } 
        } 
        return $result; 
    }

    if(defined('APP_RAN')){
        echo "HI from terminal";
    }
    if(isset($_POST['SubmitButton'])){
        $input =  $_POST['inputText'];
    }else{
        $input = readline();
    }
    $arr = json_decode($input, true);
    $output = flatten($arr);
    echo '<pre>'; print_r($output); echo '</pre>';
?>
