/*
-------------------------------------------------
   Author :       zlyuan
   date：         2019/12/22
   Description :
-------------------------------------------------
*/

package zlog2

import "testing"

func TestNew(t *testing.T) {
    conf := DefaultConfig
    l := New(conf)
    l.Info("123")
    l.Debug("123")
}
