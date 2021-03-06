import * as React from "react";
import { FunctionComponent } from "react";
import { XSidebar } from ".";
import { RouteConfig } from "../../config/routes";
import XToolbar from "./XToolbar";

type LayoutProps = {
  routeMap: RouteConfig[];
  userID?: string;
  isAdmin: boolean;
  className: string;
};

const XLayout: FunctionComponent<LayoutProps> = props => (
  <React.Fragment>
    <XToolbar />

    <XSidebar
      routeMap={props.routeMap}
      userID={props.userID}
      isAdmin={props.isAdmin}
    >
      {props.children}
    </XSidebar>
  </React.Fragment>
);

export default XLayout;
