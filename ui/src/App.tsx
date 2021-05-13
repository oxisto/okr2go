import './App.css';
import { BrowserRouter as Router, Link } from 'react-router-dom';
import { Nav, Navbar } from 'react-bootstrap';
import { Switch, Route, Redirect } from 'react-router';
import { Objectives } from './Objectives';

function App() {
  return <>
    <Router>
      <Navbar bg="light" variant="light">
        <Navbar.Brand style={{ paddingLeft: 10 }} as={Link} to="/">
          OKR
        </Navbar.Brand>
        <Nav className="mr-auto">
          <Nav.Item>
            <Nav.Link as={Link} to="/objectives" >Objectives</Nav.Link>
          </Nav.Item>
        </Nav>
      </Navbar>
      <main className="p-2">
        <Navbar></Navbar>
        <Switch>
          <Route exact path="/"><Redirect to="/objectives"></Redirect></Route>
          <Route exact path="/objectives" component={Objectives} />
        </Switch>
      </main>
    </Router>
  </>
}

export default App;
