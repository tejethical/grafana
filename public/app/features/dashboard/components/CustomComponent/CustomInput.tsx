import React from 'react';

export interface CustomInputProps {
  a: number;
  b: number;
}

export class CustomComponent extends React.Component<CustomInputProps, any> {
  constructor(props: CustomInputProps) {
    super(props);

    this.state = {
      a: 0,
      b: 0,
    };
  }

  onVariableUpdated = () => {
    this.forceUpdate();
  };

  handleChangeA = (ax: string) => {
    console.log('a changed', ax);
    this.setState({ a: ax });
  };
  handleChangeB = (bx: string) => {
    console.log('a changed', bx);
    this.setState({ b: bx });
  };

  calculateSum = () => {
    var sum = Number(this.state.a) + Number(this.state.b);
    alert('Sum is ' + sum);
  };

  render() {
    return (
      <div className="trc-component">
        <input
          value={this.state.a}
          onChange={(e) => this.handleChangeA(e.target.value)}
          type="number"
          id="n1"
          name="n1"
        />
        <input
          value={this.state.b}
          onChange={(e) => this.handleChangeB(e.target.value)}
          type="number"
          id="n2"
          name="n2"
        />
        <button onClick={(e) => this.calculateSum()}>Sum</button>
        {/* <button onClick={() => alert('Sum is ' + (this.state.a + this.state.b))}>Sum</button> */}
      </div>
    );
  }
}
