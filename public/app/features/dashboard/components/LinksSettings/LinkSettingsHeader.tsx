import React from 'react';
import { Button, Icon, HorizontalGroup } from '@grafana/ui';
import { LinkSettingsMode } from '../DashboardSettings/LinksSettings';

type LinkSettingsHeaderProps = {
  onNavClick: () => void;
  onNew: () => void;
  mode: LinkSettingsMode;
  hasLinks: boolean;
};

export const LinkSettingsHeader: React.FC<LinkSettingsHeaderProps> = ({ onNavClick, onNew, mode, hasLinks }) => {
  const isEditing = mode !== 'list';

  return (
    <div className="dashboard-settings__header">
      <HorizontalGroup align="center" justify="space-between">
        <h3>
          <span onClick={onNavClick} className={isEditing ? 'pointer' : ''}>
            Dashboard Links
          </span>
          {isEditing && (
            <span>
              <Icon name="angle-right" /> {mode === 'new' ? 'New' : 'Edit'}
            </span>
          )}
        </h3>
        {!isEditing && hasLinks ? <Button onClick={onNew}>New</Button> : null}
      </HorizontalGroup>
    </div>
  );
};
