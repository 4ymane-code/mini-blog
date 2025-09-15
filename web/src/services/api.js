import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Register from "./pages/Register";
import PostDetail from "./pages/PostDetail";

function App() {
  return (
    <Router>
      <nav style={{ padding: "1rem", borderBottom: "1px solid #ccc" }}>
        <Link to="/" style={{ marginRight: "1rem" }}>Home</Link>
        <Link to="/login" style={{ marginRight: "1rem" }}>Login</Link>
        <Link to="/register">Register</Link>
      </nav>

      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/posts/:id" element={<PostDetail />} />
      </Routes>
    </Router>
  );
}

export default App;
