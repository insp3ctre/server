import React from 'react';

import Page from '../pageFrame/Page';
import Upload from '../components/upload/Upload';

export default class UploadPage extends React.PureComponent {
    render() {
        return <Page component={Upload}/>;
    }
}