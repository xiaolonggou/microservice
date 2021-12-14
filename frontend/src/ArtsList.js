import React from 'react';
import Table from 'react-bootstrap/Table'
import axios from 'axios';

class ArtsList extends React.Component {

    readData() {
        const self = this;
        axios.get(window.global.api_location+'/arts').then(function(response) {
            console.log(response.data);

            self.setState({arts: response.data});
        }).catch(function (error){
            console.log(error);
        });
    }

    getArts() {
        let table = []

        for (let i=0; i < this.state.arts.length; i++) {

            table.push(
            <tr key={i}>
                <td>{this.state.arts[i].creator}</td>
                <td>{this.state.arts[i].description}</td>
                <td>{this.state.arts[i].format}</td>
            </tr>
            );
        }

        return table
    }

    constructor(props) {
        super(props);
        this.readData();
        this.state = {arts: []};
    
        this.readData = this.readData.bind(this);
    }

    render() {
      return (
        <div>
            <h1 style={{marginBottom: "40px"}}>Collection</h1>
            <Table>
                <thead>
                    <tr>
                        <th>
                            Author
                        </th>
                        <th>
                            Description
                        </th>
                        <th>
                            Format
                        </th>
                    </tr>
                </thead>
                <tbody>
                    {this.getArts()}
                </tbody>
            </Table>
        </div>
      ) 
    }
}

export default ArtsList;