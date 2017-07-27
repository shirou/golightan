require 'minitest/autorun'

module StringPugs
  refine String do
    def pugs
      "Pugs!"
    end
  end
end

using StringPugs

module Awesome
  def xxx
    'awesome'
  end
end

module SuperAwesome
  refine Awesome do
    def xxx
      "not #{super} but super-awesome"
    end
  end
end

using SuperAwesome

class Person
  include Awesome

  def say
    "I am #{xxx}!"
  end
end

class RefinementsTest < Minitest::Test
  # https://bugs.ruby-lang.org/issues/9451
  def test_symbol_to_proc
    assert_equal ['Pugs!', 'Pugs!'], ['a', 'b'].map(&:pugs)
  end

  # https://bugs.ruby-lang.org/issues/11476
  def test_send
    assert_equal 'Pugs!', 'a'.send(:pugs)
  end

  # https://bugs.ruby-lang.org/issues/12534
  def test_module_refinements
    person = Person.new
    assert_equal 'I am not awesome but super-awesome!', person.say
  end

  # https://bugs.ruby-lang.org/issues/7418
  def test_used_modules
    # Returns an array of all active refinements in the current scope.
    assert_equal [StringPugs, SuperAwesome], Module.used_modules
  end
end
