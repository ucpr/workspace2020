from pytest_mock import MockFixture

from main import Calc

calc = Calc()


def test_keisan():
    assert calc.keisan("hoge") != "keisan"


def test_keisan_pytest_mock(mocker: MockFixture):
    mocker.patch("main.Calc.keisan").return_value = "keisan"
    assert calc.keisan(None) == "keisan"

