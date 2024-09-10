package stingray

import (
	"github.com/gogo/protobuf/proto"
	"github.com/ngynkvn/stingray/deadlock"
)

type ModifierTableEntryHandler func(msg *deadlock.CModifierTableEntry) error

// OnModifierTableEntry registers a handler for when a ModifierBuffTableEntry
// is created or updated.
func (p *Parser) OnModifierTableEntry(fn ModifierTableEntryHandler) {
	p.modifierTableEntryHandlers = append(p.modifierTableEntryHandlers, fn)
}

// emitModifierTableEvents emits ModifierBuffTableEntry events
// from the given string table items.
func (p *Parser) emitModifierTableEvents(items []*stringTableItem) error {
	for _, item := range items {
		msg := &deadlock.CModifierTableEntry{}
		if err := proto.NewBuffer(item.Value).Unmarshal(msg); err != nil {
			_debugf("unable to unmarshal ModifierBuffTableEntry: %s", err)
			continue
		}

		for _, fn := range p.modifierTableEntryHandlers {
			if err := fn(msg); err != nil {
				return err
			}
		}
	}

	return nil
}
