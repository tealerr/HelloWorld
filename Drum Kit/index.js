    document.querySelectorAll(".drum").forEach(function (element) {
      element.addEventListener("click", function () {
        var buttonInnerHTML = this.innerHTML;
        playSound(buttonInnerHTML);
      });
    });

    document.addEventListener("keydown", function (event) {
      playSound(event.key);
    });

    function playSound(key) {
      switch (key) {
        case "w":
          var audio = new Audio("sounds/tom-1.mp3");
          audio.play();
          break;

        case "a":
          var audio = new Audio("sounds/tom-2.mp3");
          audio.play();
          break;

        case "s":
          var audio = new Audio("sounds/tom-3.mp3");
          audio.play();
          break;

        case "d":
          var audio = new Audio("sounds/tom-4.mp3");
          audio.play();
          break;

        case "j":
          var audio = new Audio("sounds/crash.mp3");
          audio.play();
          break;

        case "k":
          var audio = new Audio("sounds/kick-bass.mp3");
          audio.play();
          break;

        case "l":
          var audio = new Audio("sounds/snare.mp3");
          audio.play();
          break;

        case "p":
          var audio = new Audio("sounds/Shinunoga E-Wa.mp3");
          if (this.innerHTML==="p"){
            audio.play();
          }
          else{
            audio.pause();
          }
          break;


        default:
          console.log("Invalid key");
      }
    }