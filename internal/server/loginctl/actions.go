package loginctl

import (
	"fmt"
)

func (m *Manager) Lock() error {
	if m.sessionObj == nil {
		return fmt.Errorf("session object not available")
	}
	err := m.sessionObj.Call(dbusSessionInterface+".Lock", 0).Err
	if err != nil {
		return fmt.Errorf("failed to lock session: %w", err)
	}
	return nil
}

func (m *Manager) Unlock() error {
	err := m.sessionObj.Call(dbusSessionInterface+".Unlock", 0).Err
	if err != nil {
		return fmt.Errorf("failed to unlock session: %w", err)
	}
	return nil
}

func (m *Manager) Activate() error {
	err := m.sessionObj.Call(dbusSessionInterface+".Activate", 0).Err
	if err != nil {
		return fmt.Errorf("failed to activate session: %w", err)
	}
	return nil
}

func (m *Manager) SetIdleHint(idle bool) error {
	err := m.sessionObj.Call(dbusSessionInterface+".SetIdleHint", 0, idle).Err
	if err != nil {
		return fmt.Errorf("failed to set idle hint: %w", err)
	}
	return nil
}

func (m *Manager) Terminate() error {
	err := m.sessionObj.Call(dbusSessionInterface+".Terminate", 0).Err
	if err != nil {
		return fmt.Errorf("failed to terminate session: %w", err)
	}
	return nil
}

func (m *Manager) SetLockBeforeSuspend(enabled bool) {
	m.lockBeforeSuspend.Store(enabled)
}

func (m *Manager) SetSleepInhibitorEnabled(enabled bool) {
	m.sleepInhibitorEnabled.Store(enabled)
	if enabled {
		// Re-acquire inhibitor if enabled
		_ = m.acquireSleepInhibitor()
	} else {
		// Release inhibitor if disabled
		m.releaseSleepInhibitor()
	}
}
