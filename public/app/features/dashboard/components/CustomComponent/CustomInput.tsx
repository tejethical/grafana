// Libraries
import React, { FC, MouseEvent } from 'react';
// Components
import { IconName, IconType, IconSize } from '@grafana/ui';

interface Props {
  icon?: IconName;
  tooltip: string;
  onClick?: (event: MouseEvent<HTMLButtonElement>) => void;
  href?: string;
  children?: React.ReactNode;
  iconType?: IconType;
  iconSize?: IconSize;
}

export const CustomComponent: FC<Props> = ({}) => {
  return (
    <div className="custom-component-ust">
      <input type="number" id="n1" name="n1" />
      <input type="number" id="n2" name="n2" />
      <button>Sum</button>
    </div>
  );
};
