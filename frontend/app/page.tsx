import Header from './header';
import RequestForm from './requestForm';

export default function Home() {
    return (
      <div>
        <Header />
        <div className="grid grid-cols-2">
          <div className="p-5">
            <RequestForm />
          </div>
        </div>
      </div>
    );
  }