import addition
""" Testing program; See also: http://doc.pytest.org/en/latest/getting-started.html """


def test_add_two_numbers():
    """Tests that 'add' 2 and 5 must result in 7"""
    assert addition.add(2, 5) == 7
