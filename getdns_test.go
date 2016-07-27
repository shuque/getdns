/*
 * tests (unfinished)
 */

package getdns

import "testing"

func TestContext(t *testing.T) {
	_, err := ContextCreate()
	if err != nil {
		t.Error("contextCreate() failed:", err)
	}
}

func TestGeneral(t *testing.T) {

	ctx, err := ContextCreate()
	if err != nil {
		t.Error("ContextCreate() failed:", err)
	}

	rc, _ := General(ctx, "www.example.com", "A", nil)
	if rc != RETURN_GOOD {
		t.Error("General() failed:", rc)
	}

}
