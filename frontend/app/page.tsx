import Description from './description';
import Header from './header';
import RequestForm from './requestForm';

export default function Home() {
    return (
      <div className="h-full">
        <Header/>
        <div className="pt-20 h-full grid grid-cols-1 md:grid-cols-2">
          <div className="p-10 lg:px-20 md:self-center">
            <Description />
          </div>
          <div className="p-10 lg:px-20 md:self-center">
            <RequestForm />
          </div>
        </div>
      </div>
    );
  }