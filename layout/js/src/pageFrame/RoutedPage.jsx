import React from 'react';
import {Route, Router, browserHistory} from 'react-router';

import Player from '../components/player/Player';

import RegistrationPage from './special/RegistrationPage';

import StartpageContainer from '../containers/StartpageContainer';
import MySongsContainer from '../containers/MySongsContainer';
import MyArtistsContainer from '../containers/MyArtistsContainer';
import MyAlbumsContainer from '../containers/MyAlbumsContainer';
import MyPlaylistsContainer from '../containers/MyPlaylistsContainer';

require('./RoutedPage.scss');

export default class RoutedPage extends React.PureComponent {
    render() {
        return (
            <div className="suluvir-routed-page">
                <Router history={browserHistory}>
                    <Route component={StartpageContainer} path="/"/>

                    <Route component={RegistrationPage} path="/register"/>

                    <Route component={MySongsContainer} path="songs"/>
                    <Route component={MyAlbumsContainer} path="albums"/>
                    <Route component={MyArtistsContainer} path="artists"/>
                    <Route component={MyPlaylistsContainer} path="playlists"/>
                </Router>
                <Player />
            </div>
        );
    }
}
