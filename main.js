window.addEventListener('DOMContentLoaded', function() {
    var overlay = document.querySelector('.overlay');
    var popup = document.querySelector('.popup');
    var acceptBtn = document.querySelector('#acceptBtn');
    var refuseBtn = document.querySelector('#refuseBtn');
  
    acceptBtn.addEventListener('click', function() {
      overlay.style.display = 'none';
      popup.style.display = 'none';
      document.body.classList.remove('popup-active');
    });
  
    refuseBtn.addEventListener('click', function() {
      // Rediriger vers une page d'erreur ou afficher un message approprié
      window.location.href = 'error.html';
    });
  
    overlay.style.display = 'block';
    popup.style.display = 'block';
    document.body.classList.add('popup-active');
  });
  