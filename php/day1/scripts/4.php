<?php
    if(isset($_POST['SubmitButton4'])){
        $input =  $_POST['inputText4'];
    }else{
        $input = readline();
    }

    $obj = json_decode($input);
    // echo $obj;
    $names = array();
    for($i=0 ; $i<count($obj->{'players'});$i++){
        array_push($names,$obj->{'players'}[$i]->name);
    }
    echo "<h3>Names</h3>";
    echo '<pre>'; print_r($names); echo '</pre>';

    // array of ages
    $ages = array();
    for($i=0 ; $i<count($obj->{'players'});$i++){
        array_push($ages,$obj->{'players'}[$i]->age);
    }
    echo "<h3>Ages</h3>";
    echo '<pre>'; print_r($ages); echo '</pre>';

    // array of cities
    $cities = array();
    for($i=0 ; $i<count($obj->{'players'});$i++){
        array_push($cities,$obj->{'players'}[$i]->address->city);
    }

    echo "<h3>Cities</h3>";
    echo '<pre>'; print_r($cities); echo '</pre>';

    // print only unique names
    $unique_names = array_unique($names);

    echo "<h3>Unique Names</h3>";
    echo '<pre>'; print_r($unique_names); echo '</pre>';

    //array of names with maximum age
    $max_age = 0;
    $max_age_name = array();
    for($i=0 ; $i<count($obj->{'players'});$i++){
        $max_age = max($max_age,$obj->{'players'}[$i]->age);
    }
    for($i=0 ; $i<count($obj->{'players'});$i++){
        if($obj->{'players'}[$i]->age == $max_age)
        array_push($max_age_name,$obj->{'players'}[$i]->name);
    }
    echo "<h3>Max Age</h3>";
    echo '<pre>'; print_r($max_age); echo '</pre>';

    
?>
