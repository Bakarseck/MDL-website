window.addEventListener('DOMContentLoaded', function () {
  var overlay = document.querySelector('.overlay');
  var popup = document.querySelector('.popup');
  var acceptBtn = document.querySelector('#acceptBtn');
  var refuseBtn = document.querySelector('#refuseBtn');

  acceptBtn.addEventListener('click', function () {
    overlay.style.display = 'none';
    popup.style.display = 'none';
    document.body.classList.remove('popup-active');
  });

  refuseBtn.addEventListener('click', function () {
    // Rediriger vers une page d'erreur ou afficher un message approprié
    window.location.href = 'error.html';
  });

  overlay.style.display = 'block';
  popup.style.display = 'block';
  document.body.classList.add('popup-active');
});

const categoryList = document.querySelector('.category-list');
const categoryItems = document.querySelectorAll('.category-item');

categoryList.innerHTML += categoryList.innerHTML;

// Ajouter une classe pour définir la largeur de défilement
categoryList.classList.add('scrolling');

// Obtenir la largeur totale du défilement
let totalWidth = 0;
categoryItems.forEach(item => {
  totalWidth += item.offsetWidth + parseInt(getComputedStyle(item).marginRight);
});
categoryList.style.width = `${totalWidth}px`;
