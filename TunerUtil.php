<?php

/**
 * A utility to replace h1 and make sure we have nice buttons for more mobile friendly version
 * @author Ido Green @greenido
 * @date 3/4/2017
 */
class Ggr {

  // TODO: make sure you are working on the right input file
  private $fileName = "/Applications/MAMP/htdocs/relaymaps/index-2018.html";
  private $outputFile = "/Applications/MAMP/htdocs/relaymaps/index.html";
  private $headHTML = "/Applications/MAMP/htdocs/relaymaps/head.html";
  private $footHTML = "/Applications/MAMP/htdocs/relaymaps/footer.html";

  /**
   * Ctor
   */
  function __construct() {
  }

  function tuner() {
    $legs = 36; // we have 36 legs in the race
    $fullHtml = $this->getAllHtml();
    
    for ($i = 1; $i < $legs+1; $i++) {
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
      echo "Replacing leg $i \n";
    }
    // $headHTML = file_get_contents($this->headHTML);
    // $inx4 = strpos($fullHtml, "<head>");
    // $inx5 = strpos($fullHtml, "</head>", $inx4);
    // $fullHtml = substr_replace($fullHtml, $headHTML, $inx4, ($inx5-$inx4) );

    // $footerHtml = file_get_contents($this->footHTML);
    // $inx6 = strpos($fullHtml, "</body>");
    // $inx7 = strpos($fullHtml, "</html>", $inx6);
    // $fullHtml = substr_replace($fullHtml, $footHTML, $inx6 , $inx7 - $inx6);

    // let's output everyting to a new file
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
