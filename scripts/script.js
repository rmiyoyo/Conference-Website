const menuButton = document.querySelector('.menu-button');
const nav = document.getElementById('site-nav');

if (menuButton && nav) {
  const setOpen = (open) => {
    menuButton.setAttribute('aria-expanded', String(open));
    nav.classList.toggle('open', open);
  };

  menuButton.hidden = false;
  document.body.classList.add('has-menu');

  menuButton.addEventListener('click', () => {
    setOpen(menuButton.getAttribute('aria-expanded') !== 'true');
  });

  nav.addEventListener('click', (event) => {
    if (event.target.closest('a')) setOpen(false);
  });

  document.addEventListener('keydown', (event) => {
    if (event.key === 'Escape') setOpen(false);
  });
}

const picker = document.querySelector('.picker');

if (picker) {
  const select = picker.querySelector('select');
  const answer = picker.querySelector('.answer');

  const describe = (option) => {
    const { km, hq } = option.dataset;
    if (option.hasAttribute('data-host')) {
      return 'You are already in Nairobi. KICC is in the city centre, on Harambee Avenue.';
    }
    const border = option.hasAttribute('data-borders') ? ' It borders Nairobi.' : '';
    return `${option.textContent} is about ${km} km from KICC in a straight line. The county headquarters is ${hq}.${border}`;
  };

  picker.hidden = false;

  select.addEventListener('change', () => {
    const option = select.selectedOptions[0];
    answer.textContent = option && option.value ? describe(option) : '';
  });
}
