import { useEffect } from 'react';
import { useRouter } from 'next/navigation';

const withAuth = (WrappedComponent) => {
  return (props) => {
    const router = useRouter();
    useEffect(() => {
      const authToken = localStorage.getItem('authToken');
      if (!authToken) {
        router.push('/auth');
      }
    }, [router]);

    return <WrappedComponent {...props} />;
  };
};

export default withAuth;
