String.prototype.format = function() {
    return [...arguments].reduce((p,c) => p.replace(/%s/,c), this);
  };
var data = [
    {
        href: '#job-abcc',
        date: '03/2022 - Present',
        title: 'Analyst Programmer in ABCC.com',
        icon: '<i class="fa-brands fa-ethereum"></i>  <i class="fa-solid fa-bitcoin-sign" style="color: #f78f15;"></i>',
        detail: 'With nice colleages, I work in Backend development of a cryptocurrency exchange platform, ABCC.com. Learnt a lots to code with organized structure and high reliability'
    },
    {
        href: '#job-hkpc',
        date: '07/2021 - 02/2022',
        title: 'Engineering Officer in HKPC.org',
        icon: '<i class="fa-solid fa-microchip" style="color:#ff70a6"></i>',
        detail: 'My starting point of carreer and I worked in Smart City Division in HKPC. The company provided lots of oppunitry to work in smart city application, with government and various external parties.'
    },
    {
        href: '#study-cityu',
        date: '09/2018 - 07/2021',
        title: 'BEng in City University of Hong Kong',
        icon: '<i class="fa-solid fa-graduation-cap" style="color:#ff9770"></i>',
        detail: 'I was a Engineering student majored in Electronic and Communication Engineering and I got nearly first honours but second up. The degree highly train up my critical thinking, and I met the repectful professor that give me oppunity to work on cyber security in IoT applications.'
    },
    {
        href: '#study-hku',
        date: '09/2016 - 07/2018',
        title: 'AEng in HKU SPACE',
        icon: '<i class="fa-sharp fa-solid fa-school" style="color:#f5c510;"></i>',
        detail: 'I play bass and guitar, why not study more about these! With an interest about audio, I studied an Eletronic Engineering, after I graduated from secondary school.'
    },
    {
        href: '',
        date: '1998',
        title: 'Born in HK',
        icon: '<i class="fa-sharp fa-solid fa-fire" style="color:red;"></i>',
        detail: 'Thank you my family'
    }
]

var contentTemplate = `\
<div class="timeline">
    <a href="/blog/about%s" class="timeline-content">
        <div class="timeline-icon">%s</div>
        <h3 class="title">%s</h3>
        <h3 class="title">%s</h3>
        <p class="description">%s</p>
    </a>
</div>
`

$(document).ready(function(){
    data.forEach(d =>$('#roadmap').append(contentTemplate.format(
        d.href,d.icon,d.date,d.title,d.detail
    )))

    $('#my-roadmap').roadmap(events, {
        prevArrow: '<i class="fa-solid fa-arrow-left"></i>',
        nextArrow: '<i class="fa-solid fa-arrow-right"></i>',
    });
});