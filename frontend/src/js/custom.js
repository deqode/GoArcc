 $('.palceholder').click(function() {
        $(this).siblings('input').focus();
      });       
      $('div.field-group input.field-control').focus(function() {
        
        $(this).parent().addClass ("focused")
      });

      $('.field-control').blur(function() {
        var $this = $(this);
        if ($this.val().length == 0)
          $(this).parent().removeClass("focused");
      });
      $('.field-control').blur();

    
$(document).ready(function() {
       $(".app-header__menu").click(function () {
          toggleButton = $(".app-container")
          if ($(toggleButton).hasClass("header-menu-open")) {
             $(".app-container").removeClass("header-menu-open");
          }
          else {
             $(".app-container").addClass("header-menu-open"); // instead of this do the below 
          } 
       });
    });



 var removeClass = true;
        $(".mobile-toggle-nav").click(function () {
            $(".app-container").toggleClass('sidebar-open');
            removeClass = false;
        });

        $(".app-container").click(function() {
            removeClass = false;
        });

        $(".app-container").click(function () {
            if (removeClass) {
                $(".app-container").removeClass('sidebar-open');
            }
            removeClass = true;
        });
