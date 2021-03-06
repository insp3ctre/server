import React from 'react';
import Immutale from 'immutable';
import {connect} from 'react-redux';

import PlaylistList from './PlaylistList';
import Loading from '../util/Loading';

import {fetchMyPlaylists} from '../../actions/thunkActions';

class MyPlaylistList extends React.PureComponent {
    componentWillMount() {
        const {fetchMyPlaylists} = this.props;
        fetchMyPlaylists();
    }

    render() {
        const {playlists} = this.props;

        if (playlists === null) {
            return <Loading />;
        }

        return <PlaylistList playlists={playlists} />;
    }
}

MyPlaylistList.propTypes = {
    fetchMyPlaylists: React.PropTypes.func.isRequired,
    playlists: React.PropTypes.instanceOf(Immutale.List).isRequired
}

function mapStateToProps(state) {
    return {
        playlists: state.myPlaylists
    }
}

export default connect(mapStateToProps, {fetchMyPlaylists})(MyPlaylistList);