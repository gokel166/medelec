import React from 'react';
import ReactDom from 'react-dom';
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';

import Home from './home';
import Other from './other';
import Topic from './topic'

ReactDOM.render(
  <BrowserRouter>
    <div>
        <ul>
            <li><Link to="/">Home</Link></li>
            <li><Link to="/other">Other</Link></li>
            <li><Link to="/topics">Topics</Link></li>
        </ul>

        <hr/>

        <Route exact path="/" component={Home}/>
        <Route path="/other" component={Other}/>
        <Route path="/topics" component={Topics}/>
    </div>
  </BrowserRouter>
  document.getElementById('app')
);
