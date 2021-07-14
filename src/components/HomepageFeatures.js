import React from 'react';
import clsx from 'clsx';
import styles from './HomepageFeatures.module.css';

const FeatureList = [
  {
    title: 'Rich Support',
    Svg: require('../../static/img/logo.svg').default,
    description: (
      <>
        GoArcc supports gRPC, REST, GraphQl. 
        We have covered everything you need to start writing services.
      </>
    ),
  },
  {
    title: 'Focus on What Matters',
    Svg: require('../../static/img/logo.svg').default,
    description: (
      <>
        Start writing the business logic part of your
        services and deploy your code within a minute.
      </>
    ),
  },
  {
    title: 'Powered by Deqode',
    Svg: require('../../static/img/undraw_docusaurus_react.svg').default,
    description: (
      <>
        Extend or customize your website layout by reusing React. Docusaurus can
        be extended while reusing the same header and footer.
      </>
    ),
  },
];

function Feature({Svg, title, description}) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} alt={title} />
      </div>
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
