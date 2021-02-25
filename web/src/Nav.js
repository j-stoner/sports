import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';
import Nav from 'react-bootstrap/Nav';
import { Link } from 'react-router-dom';
import { withRouter } from "react-router";

class TopNav extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <Container fluid>
                <Navbar expand="sm" bg="dark" variant="dark" hidden={this.props.location.pathname === "/board" ? true : false}>
                    <Navbar.Toggle aria-controls="basic-navbar-nav"></Navbar.Toggle>
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Nav className="mr-auto">
                            <Nav.Link as={Link} to="/">Home</Nav.Link>
                            <Nav.Link as={Link} to="/nhl">NHL</Nav.Link>
                            <Nav.Link as={Link} to="/mlb">MLB</Nav.Link>
                            <Nav.Link as={Link} to="/ncaam">NCAA Men Basketball</Nav.Link>
                            <Nav.Link as={Link} to="/img">Image Board</Nav.Link>
                            <Nav.Link as={Link} to="/clock">Clock</Nav.Link>
                            <Nav.Link as={Link} to="/sys">System Info</Nav.Link>
                            <Nav.Link as={Link} to="/board">Live Board</Nav.Link>
                        </Nav>
                    </Navbar.Collapse>
                </Navbar>
            </Container >
        );
    }
}

export default withRouter(TopNav);