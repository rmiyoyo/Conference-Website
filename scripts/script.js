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
    username: 'Naledi Moloi',
    occupation: 'Artificial Intelligence Researcher, Johannesburg AI',
    jd: "Naledi's research focuses on developing AI applications for solving African challenges in areas such as healthcare, agriculture, and education. She is also an advocate for responsible AI development and is involved in several initiatives to promote AI ethics.",
  },
  {
    id: 1,
    profileID: './images/keynote/lwazi.png',
    username: 'Lwazi Mhlongo',
    occupation: 'Software Developer, Durban Gram',
    jd: 'Lwazi is the founder of a tech community that promotes diversity and inclusion in the industry. He is also a skilled software developer with expertise in web development and cloud computing.',
  },
  {
    id: 2,
    profileID: './images/keynote/yemisi.png',
    username: 'Yemisi Akindele',
    occupation: 'Digital Marketing Manager, Lagos PR',
    jd: 'Yemisi has over a decade of experience in digital marketing and has led successful campaigns for several brands across Africa. She is also a thought leader in the field and has been featured in several industry publications.',
  },
  {
    id: 3,
    profileID: './images/keynote/hicham.png',
    username: 'Hicham Cherkaoui',
    occupation: 'Information Security Consultant, Casablanca Defense',
    jd: 'Hicham has over 15 years of experience in information security and has worked with several organizations across Africa and Europe. He is also a frequent speaker at security conferences and events.',
  },
  {
    id: 4,
    profileID: './images/keynote/doreen.png',
    username: 'Doreen Macharia',
    occupation: 'User Experience (UX) Designer, Gearbox',
    jd: 'Doreen is passionate about creating user-centered digital experiences that solve real-world problems. She has worked on several projects that have won awards for their innovative UX design.',
  },
  {
    id: 5,
    profileID: './images/keynote/fadumo.png',
    username: 'Fadumo Mohamed',
    occupation: 'Data Analyst, Mogadishu Health Systems',
    jd: 'Fadumo specializes in data analysis and visualization and is responsible for analyzing healthcare data to improve patient outcomes. She is also a mentor to young women interested in pursuing careers in STEM fields.',
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
