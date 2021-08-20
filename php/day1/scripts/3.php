<?php
    function CamelCase($arr){
        $ans=array();
        foreach($arr as $i){
            $curr = explode("_",$i);
            $newString ="";
            $newString.=$curr[0];
            for($j=1;$j<count($curr);$j++){
                $newString.=ucwords($curr[$j]);
            }
            array_push($ans,$newString);
        }
        return $ans;
    }

    if(isset($_POST['SubmitButton3'])){
        $input =  $_POST['inputText3'];
    }else{
        $input = readline();
    }
    $arr = json_decode($input, true);
    $output = CamelCase($arr);
    echo '<pre>'; print_r($output); echo '</pre>';
?>
