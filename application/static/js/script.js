var card = document.getElementById("bc-front")
var card_inner = document.getElementById("card-inner")
var flipped = false;

function Rotate() {
  if (flipped === false){
    card_inner.setAttribute("style", "transform: rotateY(-180deg)");
    flipped = true;
  } else if (flipped === true) {
    card_inner.setAttribute("style", "transform: rotateY(0deg)");
    flipped = false;
  }

}
