package sqlparser

import (
	"testing"

	"github.com/zssky/log"
)

func TestStatement(t *testing.T)  {
	trigger1()

	trigger2()
	//
	//
	trigger3()
	//
	//global()
}

func trigger1() {
	//
	sql := "CREATE DEFINER=`root`@`localhost` TRIGGER `pt_osc_test_test02_ins` AFTER INSERT ON `test`.`test02` FOR EACH ROW REPLACE INTO `test`.`_test02_new` (`id`, `name`, `address`) VALUES (NEW.`id`, NEW.`name`, NEW.`address`)"
	stmt, err := Parse(sql)
	if err != nil {
		log.Fatal(err)
	}

	switch st := stmt.(type) {
	case *DDL:
		log.Infof("table %v", st.Table)
		switch st.Action {
		case CreateTrggerStr:
			log.Infof("create trigger %v", st.Trigger)
			log.Infof("definer %v", st.Definer)
			log.Infof("body %s", st.Trigger.Body)
		}
	case *DBDDL:
		log.Infof(st.DBName)
	}
}

func trigger3() {
	//
	sql := "CREATE DEFINER=`root`@`localhost` TRIGGER `pt_osc_test_test02_del` AFTER DELETE ON `test`.`test02` FOR EACH ROW DELETE IGNORE FROM `test`.`_test02_new` WHERE `test`.`_test02_new`.`id` <=> OLD.`id`"
	stmt, err := Parse(sql)
	if err != nil {
		log.Fatal(err)
	}

	switch st := stmt.(type) {
	case *DDL:
		log.Infof("table %v", st.Table)
		switch st.Action {
		case CreateTrggerStr:
			log.Infof("create trigger %v", st.Trigger)
			log.Infof("definer %v", st.Definer)
			log.Infof("body %s", st.Trigger.Body)
		}
	case *DBDDL:
		log.Infof(st.DBName)
	}
}


func trigger2() {
	//
	sql := "CREATE DEFINER=`root`@`localhost` TRIGGER `pt_osc_test_test02_upd` AFTER UPDATE ON `test`.`test02` FOR EACH ROW BEGIN DELETE IGNORE FROM `test`.`_test02_new` WHERE !(OLD.`id` <=> NEW.`id`) AND `test`.`_test02_new`.`id` <=> OLD.`id`;REPLACE INTO `test`.`_test02_new` (`id`, `name`, `school`) VALUES (NEW.`id`, NEW.`name`, NEW.`school`);END"
	stmt, err := Parse(sql)
	if err != nil {
		log.Fatal(err)
	}

	switch st := stmt.(type) {
	case *DDL:
		log.Infof("table %v", st.Table)
		switch st.Action {
		case CreateTrggerStr:
			log.Infof("create trigger %v", st.Trigger)
			log.Infof("definer %v", st.Definer)
			log.Infof("body %s", st.Trigger.Body)
		}
	case *DBDDL:
		log.Infof(st.DBName)
	}
}

func global()  {
	sql := "set @@session.gtid_mode=on # get master status"
	stmt, err := Parse(sql)
	if err != nil {
		log.Fatal(err)
	}

	switch st := stmt.(type) {
	case *DDL:
		log.Infof("table %v", st.Table)
	case *DBDDL:
		log.Infof(st.DBName)
	case *Set:
		log.Infof("comments %v", st)
	}
}