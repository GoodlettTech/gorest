import './App.css';
import AuthPage from './pages/Auth';
import { getToken } from './services/Auth';

function App() { 
  return (
    <div className="container-lg">
      <div className="flex flex-row">
        <div>
          {getToken()}
        </div>
      </div>
      <AuthPage/>
    </div>
  );
}

export default App;
