<?php
    function masked_number($phone_number){
        $masked_num="";
            for($i=0;$i<strlen($phone_number);$i++){
            if($i>=2 && $i<=7){
                $masked_num.="*";
            }
            else{
                $masked_num.=$phone_number[$i];
            }
        }
        return $masked_num;
    }

    if(isset($_POST['SubmitButton2'])){
        $input =  $_POST['inputText2'];
    }else{
        $input = readline();
    }
    $output = masked_number($input);
    echo '<pre>'; print_r($output); echo '</pre>';
?>
