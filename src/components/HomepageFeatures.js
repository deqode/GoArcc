import React from 'react';
import clsx from 'clsx';
import styles from './HomepageFeatures.module.css';

const FeatureList = [
  {
    title: 'Rich Support',
    Svg: require('../../static/img/rich_support.svg').default,
    description: (
      <>
        GoArcc supports gRPC, REST, GraphQl. 
        We have covered everything you need to start writing your services.
      </>
    ),
  },
  {
    title: 'Focus on What Matters',
    Svg: require('../../static/img/focus.svg').default,
    description: (
      <>
        Start writing the business logic part of your
        services and deploy your code within a minute.
      </>
    ),
  },
  {
    title: 'Powered by Deqode',
    Svg: require('../../static/img/deq.svg').default,
    description: (
      <>
        Constanly updated and maintained by developers at Deqode.
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
