const menu = document.getElementById('toggle');
const closeMenu = document.querySelector('.nav-links i');
const showMenu = document.querySelector('.toggle-menu');
const speakers = document.querySelector('.guests');
const navItems = document.querySelectorAll('.toggle-menu li a');

menu.addEventListener('click', () => {
  showMenu.style.display = 'absolute';
  showMenu.style.left = '0';
});

closeMenu.addEventListener('click', () => {
  showMenu.style.display = 'absolute';
  showMenu.style.left = '-100%';
});

const speakersData = [
  {
    id: 0,
    profileID: './images/keynote/naledi.png',
    username: 'James Smith',
    occupation: 'Chief Technology Officer, Innovate Tech',
    jd: "James leads a team of engineers and researchers focused on developing cutting-edge technology solutions. He has over 10 years of experience in software development and is passionate about leveraging technology to drive innovation.",
  },
  {
    id: 1,
    profileID: './images/keynote/lwazi.png',
    username: 'Sarah Johnson',
    occupation: 'Cybersecurity Expert, SecureNet',
    jd: 'With a background in computer science and extensive experience in cybersecurity, Sarah works to protect organizations from cyber threats. She frequently speaks at industry conferences about best practices in security.',
  },
  {
    id: 2,
    profileID: './images/keynote/yemisi.png',
    username: 'Michael Lee',
    occupation: 'Data Scientist, DataWise',
    jd: 'Michael specializes in machine learning and data analytics. He has worked on various projects that utilize data to improve business outcomes and is an advocate for data-driven decision making.',
  },
  {
    id: 3,
    profileID: './images/keynote/hicham.png',
    username: 'Emily Chen',
    occupation: 'Product Manager, Tech Innovations',
    jd: 'Emily has a strong background in product management and user experience design. She is dedicated to creating products that meet user needs and enhance customer satisfaction.',
  },
  {
    id: 4,
    profileID: './images/keynote/doreen.png',
    username: 'Robert Brown',
    occupation: 'Blockchain Developer, Crypto Solutions',
    jd: 'Robert is an expert in blockchain technology and its applications. He has worked on several high-profile projects that utilize blockchain for secure transactions and data integrity.',
  },
  {
    id: 5,
    profileID: './images/keynote/fadumo.png',
    username: 'Linda Garcia',
    occupation: 'AI Researcher, Future AI Labs',
    jd: 'Linda focuses on developing artificial intelligence algorithms that can solve complex problems. She is passionate about ethical AI development and participates in various initiatives promoting responsible AI practices.',
  },
];

speakers.innerHTML = speakersData.map((talkInfo) => `
<div class="talk-details">
  <img src="${talkInfo.profileID}" alt="Speakers"/>
  <div class="talk-info">
    <h3 class="talk-id">${talkInfo.username}</h3>
    <p class="talk-geo">${talkInfo.occupation}</p>
    <p class="talk-awrds">${talkInfo.jd}</p>
  </div>
</div>`).join('');

navItems.forEach((item) => {
  item.addEventListener('click', () => {
    showMenu.style.display = 'absolute';
    showMenu.style.left = '-100%';
  });
});
