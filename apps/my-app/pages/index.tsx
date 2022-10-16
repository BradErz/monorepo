import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";
import styles from "../styles/Home.module.css";
import { ReviewsService } from "@buf/bufbuild_connect-web_braderz_monorepo/reviews/v1/reviews_connectweb";
import {
  ListReviewsRequest,
  ListReviewsResponse,
} from "@buf/bufbuild_connect-web_braderz_monorepo/reviews/v1/reviews_pb";
import {
  createPromiseClient,
  createConnectTransport,
} from "@bufbuild/connect-web";
import { useState } from "react";

interface Review {
  productId: string;
  title: string;
  body: string;
  rating: number;
}

const Home: NextPage = () => {
  const [reviews, setReviews] = useState<Review[]>([]);

  const reviewsClient = createPromiseClient(
    ReviewsService,
    createConnectTransport({
      baseUrl: "http://localhost:10000",
    })
  );

  const fetchReviews = async () => {
    setReviews([])
    const request = new ListReviewsRequest({
      productId: "6341bb43a987d8ca2473f21e",
    });
    const response = await reviewsClient.listReviews(request);

    // update the state
    for (const review of response.reviews) {
      setReviews((resp) => [
        ...resp,
        {
          productId: review.productId,
          title: review.title,
          body: review.body,
          rating: review.rating,
        },
      ]);
    }
  };

  return (
    <div className={styles.container}>
      <main className={styles.main}>
        <h1 className={styles.title}>
          Welcome to <a href="https://nextjs.org">Next.js!</a>
        </h1>

        <button className={styles.button} onClick={fetchReviews}>
          fetchReviews
        </button>

        <div className={styles.container}>
          {reviews.map((resp, i) => {
            return (
              <div key={`resp${i}`} className={styles.card}>
                <p className={styles.respText}>Title: {resp.title}</p>
                <p className={styles.respText}>Body: {resp.body}</p>
                <p className={styles.respText}>Rating: {resp.rating}</p>
                <p className={styles.respText}>Product ID: {resp.productId}</p>
              </div>
            );
          })}
        </div>
      </main>

      <footer className={styles.footer}>
        <a
          href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
          target="_blank"
          rel="noopener noreferrer"
        >
          Powered by{" "}
          <span className={styles.logo}>
            <Image src="/vercel.svg" alt="Vercel Logo" width={72} height={16} />
          </span>
        </a>
      </footer>
    </div>
  );
};

export default Home;
