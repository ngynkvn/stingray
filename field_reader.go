package stingray

import (
	"strings"
)

func readFields(r *reader, s *serializer, state *fieldState) {
	fps := readFieldPaths(r)

	for _, fp := range fps {
		decoder := s.getDecoderForFieldPath(fp, 0)

		// TODO: Refactor out logging blocks
		name := strings.Join(s.getNameForFieldPath(fp, 0), ".")
		typ := s.getTypeForFieldPath(fp, 0)
		field := s.getFieldForFieldPath(fp, 0)
		Dbg.Debug("Reading field",
			"position", r.position(),
			"path", fp.String(),
			"name", name,
			"serializer", s.name,
			"type", typ,
			"decoder", _nameof(decoder),
			"encoder", field.encodeFlags,
			"model", field.modelString())
		dump.With(
			"pos", r.position(),
			"path", fp.String(),
			"name", name,
			"type", typ,
			"decoder", _nameof(decoder),
			"ser", s.name,
			"encoder", field.encoder).Info("")

		val := decoder(r)
		state.set(fp, val)

		// TODO: Refactor out logging blocks
		if Level.Level() == DEBUG {
			name := strings.Join(s.getNameForFieldPath(fp, 0), ".")
			fp2 := newFieldPath()
			b := s.getFieldPathForName(fp2, name)

			if !b {
				_panicf("GOT NO FP: name=%s fp2=%#vv", name, fp2)
			}

			if fp2.String() != fp.String() {
				_panicf("GOT FP MISMATCH: fp=%s fp2=%s", fp, fp2)
			}

			fp2.release()

			Dbg.Debug("Field value", "value", val)
		}

		fp.release()
	}
}
