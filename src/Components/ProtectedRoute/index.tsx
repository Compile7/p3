import React from 'react';
import { Navigate, RouteProps, useLocation } from 'react-router-dom';
import { isLoggedIn } from '../../utils/index';

export const ProtectedRoute: React.FC<RouteProps> = (props) => {
    const { pathname, search } = useLocation();
    const { children} = props;
    if (!isLoggedIn()) {
        
        return search ? (
            <Navigate to={`/?redirect=${pathname}&${search.slice(1)} `} replace/>
            
        ) : (
            <Navigate to={`/?redirect=${pathname}`} replace/>
           
        );
    } else
        return children
};
