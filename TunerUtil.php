<?php

/**
 * A utility to replace h1 and make sure we have nice buttons for more mobile friendly version
 * @author Ido Green @greenido
 * @date 3/4/2017
 */
class Ggr {

  private $fileName = "/Applications/MAMP/htdocs/relaymaps/index.html";
  private $outputFile = "/Applications/MAMP/htdocs/relaymaps/index-o.html";
  
  /**
   * Ctor
   */
  function __construct() {
  }

  function tuner() {
    $legs = 36; // we have 36 legs in the race
    $fullHtml = $this->getAllHtml();
    
    for ($i = 2; $i < $legs+1; $i++) {
      $inx1 = strpos($fullHtml, "<h1>Leg") - 7;
      if ($i > 2) {
        $fullHtml = substr_replace($fullHtml, '</div></div> <button class="btn btn-primary btn-space" type="button" data-toggle="collapse" 
            data-target="#leg-' . $i . '-details" aria-expanded="false" aria-controls="leg-' . $i . '-details">
          Leg '.$i.'</button> <div class="collapse" id="leg-' . $i .'-details">
            <div class="card card-block"> ', $inx1, 15);
        
      } 
      else {
        $fullHtml = substr_replace($fullHtml, '<button class="btn btn-primary btn-space" type="button" data-toggle="collapse" 
            data-target="#leg-' . $i . '-details" aria-expanded="false" aria-controls="leg-' . $i . '-details">
          Leg '.$i.'</button> <div class="collapse" id="leg-' . $i .'-details">
            <div class="card card-block">', $inx1, 15); 
      } 
    }
    file_put_contents($this->outputFile, $fullHtml);
  }
  
  function getAllHtml() {
    $fullText =  file_get_contents($this->fileName);
    return $fullText;
  }
}

//
// Start the party
// 
$analyzer = New Ggr();
$analyzer->tuner();
