//import index from 'axios';

const axios = require('axios');
const cheerio = require('cheerio');

//base URL used for most requests
const baseUrl = 'http://www.cruchroll.com/';

//main module
const Crunchyroll = {
  async getAllSeries() {
    //load catalog
    const {data} = await axios.get(`${baseUrl}/videos/anime`);
    // create cheerio cursor
    const $ = cheerio.load(data);

    const content = $('div.main_content');
    const series = $('li.group-item');
    series.each((index, el) => {
      console.log($(el).html());
    });
    console.log(data);
  },
  getEpisodes(series) {
    //
  },
  getMySeries() {
    //
  },
  search(query) {
    //
    console.log(query)
  },
};

Crunchyroll.getAllSeries();

module.exports = Crunchyroll;
